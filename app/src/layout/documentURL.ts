export const documentURL = (href: string, id: string) => {
    const url = new URL(href);
    url.searchParams.delete("url");
    url.searchParams.delete("focus");
    url.searchParams.delete("fullscreen");
    url.searchParams.delete("id");
    if (id) {
        url.searchParams.set("id", id);
    }
    return url.href;
};
