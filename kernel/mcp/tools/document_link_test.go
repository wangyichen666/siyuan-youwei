// SiYuan - From thought to insight, with agents
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

package tools

import "testing"

func TestResolveDocumentLinkID(t *testing.T) {
	const id = "20260927000000-abcdefg"
	for _, input := range []string{
		id,
		"http://127.0.0.1:6806/stage/build/desktop/?r=kanv0uc&id=" + id,
		"http://127.0.0.1:6806/?id=" + id,
		"siyuan://blocks/" + id,
		"web+siyuan://blocks/" + id,
	} {
		resolved, err := resolveDocumentLinkID(input)
		if err != nil || resolved != id {
			t.Errorf("resolveDocumentLinkID(%q) = %q, %v", input, resolved, err)
		}
	}
	for _, input := range []string{
		"http://127.0.0.1:6806/stage/build/desktop/?r=kanv0uc",
		"http://127.0.0.1:6806/other/?id=" + id,
		"file:///stage/build/desktop/?id=" + id,
		"siyuan://blocks/not-a-block-id",
	} {
		if _, err := resolveDocumentLinkID(input); err == nil {
			t.Errorf("resolveDocumentLinkID(%q) accepted an invalid document link", input)
		}
	}
}

func TestNormalizeDocumentToolURL(t *testing.T) {
	const id = "20260927000000-abcdefg"
	args := map[string]any{"url": "http://127.0.0.1:6806/?id=" + id}
	if err := normalizeDocumentToolURL(args, "parentID", "parentID"); err != nil || args["parentID"] != id {
		t.Fatalf("normalizeDocumentToolURL() = %#v, %v", args, err)
	}
	args = map[string]any{"id": id, "url": "siyuan://blocks/20260927000000-otherid"}
	if err := normalizeDocumentToolURL(args, "id", "id"); err == nil {
		t.Fatal("conflicting id and url were accepted")
	}
}
