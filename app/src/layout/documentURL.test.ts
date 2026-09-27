import {strict as assert} from "node:assert";
import {describe, it} from "node:test";
import {documentURL} from "./documentURL";

describe("documentURL", () => {
    it("preserves unrelated query parameters and replaces a stale document target", () => {
        const href = "http://127.0.0.1:6806/stage/build/desktop/?r=kanv0uc&url=siyuan%3A%2F%2Fblocks%2Fold&focus=1&id=old";
        const result = new URL(documentURL(href, "20260927000000-abcdefg"));
        assert.equal(result.searchParams.get("id"), "20260927000000-abcdefg");
        assert.equal(result.searchParams.get("r"), "kanv0uc");
        assert.equal(result.searchParams.has("url"), false);
        assert.equal(result.searchParams.has("focus"), false);
        assert.ok(result.href.endsWith("id=20260927000000-abcdefg"));
    });

    it("removes the document target when no document is active", () => {
        const result = new URL(documentURL("http://127.0.0.1:6806/stage/build/desktop/?r=kanv0uc&id=old", ""));
        assert.equal(result.searchParams.has("id"), false);
        assert.equal(result.searchParams.get("r"), "kanv0uc");
    });
});
