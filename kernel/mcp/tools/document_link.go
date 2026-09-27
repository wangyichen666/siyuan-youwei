// SiYuan - From thought to insight, with agents
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

package tools

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/siyuan-note/siyuan/kernel/util"
)

var documentLinkID = regexp.MustCompile(`^\d{14}-\w{7}$`)

// resolveDocumentLinkID 从块链接提取 ID，不访问链接指向的服务器。
func resolveDocumentLinkID(value string) (string, error) {
	value = strings.TrimSpace(value)
	if !strings.Contains(value, "://") {
		return value, nil
	}
	parsed, err := url.Parse(value)
	if err != nil {
		return "", fmt.Errorf("invalid document link: %w", err)
	}
	var id string
	switch parsed.Scheme {
	case "http", "https":
		if parsed.Host == "" || (parsed.Path != "/" && parsed.Path != "" && parsed.Path != "/stage/build/desktop/") {
			return "", fmt.Errorf("unsupported document link: %s", value)
		}
		id = parsed.Query().Get("id")
	case "siyuan", "web+siyuan":
		if parsed.Host != "blocks" {
			return "", fmt.Errorf("unsupported document link: %s", value)
		}
		id = strings.TrimPrefix(parsed.Path, "/")
	default:
		return "", fmt.Errorf("unsupported document link: %s", value)
	}
	if !documentLinkID.MatchString(id) {
		return "", fmt.Errorf("document link must contain a valid block ID: %s", value)
	}
	return id, nil
}

func normalizeDocumentLinkArgs(args map[string]any, fields ...string) error {
	for _, field := range fields {
		value, ok := args[field].(string)
		if !ok || value == "" {
			continue
		}
		id, err := resolveDocumentLinkID(value)
		if err != nil {
			return err
		}
		args[field] = id
	}
	return nil
}

func normalizeDocumentToolURL(args map[string]any, target string, fields ...string) error {
	if rawURL, ok := args["url"].(string); ok && rawURL != "" {
		id, err := resolveDocumentLinkID(rawURL)
		if err != nil {
			return err
		}
		if existing, ok := args[target].(string); ok && existing != "" {
			existingID, resolveErr := resolveDocumentLinkID(existing)
			if resolveErr != nil {
				return resolveErr
			}
			if existingID != id {
				return fmt.Errorf("id and url refer to different blocks")
			}
		}
		args[target] = id
	}
	return normalizeDocumentLinkArgs(args, fields...)
}

func documentWebURL(id string) string {
	return "http://127.0.0.1:" + util.ServerPort + "/stage/build/desktop/?id=" + url.QueryEscape(id)
}
