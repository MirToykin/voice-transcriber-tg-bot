package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MirToykin/voice-transcriber-tg-bot/lib/e"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	getFile           = "getFile"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host         string
	basePath     string
	baseFilePath string
	client       http.Client
}

func New(host string, token string) Client {
	basePath := newBasePath(token)

	return Client{
		host:         host,
		basePath:     basePath,
		baseFilePath: newBaseFilePath(basePath),
		client:       http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func newBaseFilePath(basePath string) string {
	return path.Join("file", basePath)
}

func (c *Client) FileFullPath(filePath string) string {
	return "https://" + path.Join(c.host, c.baseFilePath, filePath)
}

func (c *Client) Updates(ctx context.Context, offset, limit int) (updates []Update, err error) {
	defer func() { err = e.WrapIfErr("failed to get updates", err) }()

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(ctx, getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessage(ctx context.Context, chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(ctx, sendMessageMethod, q)
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}

	return nil
}

func (c *Client) SendReplyMessage(ctx context.Context, chatID int, text string, messageID int) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)
	q.Add("reply_to_message_id", strconv.Itoa(messageID))

	_, err := c.doRequest(ctx, sendMessageMethod, q)
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}

	return nil
}

func (c *Client) File(ctx context.Context, fileId string) (file *File, err error) {
	defer func() { err = e.WrapIfErr("failed to get file", err) }()
	q := url.Values{}
	q.Add("file_id", fileId)

	data, err := c.doRequest(ctx, getFile, q)
	if err != nil {
		return nil, err
	}

	var fileRes FileResponse
	err = json.Unmarshal(data, &fileRes)
	if err != nil {
		return nil, err
	}

	if !fileRes.OK {
		return nil, errors.New(fmt.Sprintf("code: %d, description: %s", fileRes.ErrorCode, fileRes.Description))
	}

	file = &fileRes.Result

	return file, nil
}

func (c *Client) doRequest(ctx context.Context, method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("failed to do request", err) }()

	reqUrl := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = res.Body.Close() }()

	return io.ReadAll(res.Body)
}
