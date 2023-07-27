(() => {
    if (navigator.clipboard === undefined) {
        console.warn("clipboard is not support!")
        return
    }
    let lang = navigator.language || navigator.userLanguage;//常规浏览器语言和IE浏览器
    lang = lang.substr(0, 2);
    let shareBtn = document.createElement("div")
    switch (lang) {
        case "zh":
            shareBtn.innerText = "分享"
            break;
        default:
            shareBtn.innerText = "SHARE"
            break;
    }

    shareBtn.className = "share-button"
    window.addEventListener("load", () => {
        document.getElementById("note-body").appendChild(shareBtn)
        shareBtn.addEventListener("click", () => {
            let shareData = ""
            switch (lang) {
                case "zh":
                    shareData = `分享了一篇笔记：${document.title}， ${document.URL}`
                    break;
                default:
                    shareData = `Share a note: ${document.title}, ${document.URL}`
                    break;
            }
            navigator.clipboard.writeText(shareData);
        })
    })

})()
