let popelement = null
let hoverTriggerElement
let trigger = {
    id: 0
}
function onDoclinkHover(node, url, p) {
    hoverTriggerElement = node
    trigger.id += 1
    let id = trigger.id
    setTimeout(() => {
        if (hoverTriggerElement === node && id === trigger.id) {
            if (isRootWindow()) {
                createIframe(getNodePos(node), url, p)
            } else {
                window.parent.postMessage({ node: getNodePos(node), url, p }, "*")
            }
        }

    }, 500);

}

function getNodePos(node) {
    let pos = node.getBoundingClientRect();
    return {
        offsetLeft: node.offsetLeft,
        offsetWidth: node.offsetWidth,
        offsetTop: node.offsetTop,
        offsetHeight: node.offsetHeight,
        bottom: pos.bottom,
        right: pos.right,
    }
}
let lastNode = null
function createIframe(node, url, p) {
    if (node != null) { lastNode = node }
    else { node = lastNode }
    if (lastDocOut != -1) clearTimeout(lastDocOut)
    lastDocOut = -1
    closeIframe()
    popelement = document.createElement("iframe")
    popelement.setAttribute("src", url)
    popelement.setAttribute("frameBorder", 0)
    popelement.className = "double-chain-iframe"
    let popW = 600
    let popH = 400
    popelement.style.width = `${popW}px`
    popelement.style.height = `${popH}px`
    popelement.style.position = "absolute"

    let screenH = window.innerHeight
    let screenW = window.innerWidth


    let linkBottom = screenH - node.bottom
    let linkRight = screenW - node.right

    let left = node.offsetLeft + node.offsetWidth
    let top = node.offsetTop + node.offsetHeight
    let right = null
    if (linkRight < popW) {
        left = null
        right = "10px"
    }

    if (linkBottom < popH) {
        top = node.offsetTop - popH
    }

    if (left != null) popelement.style.left = `${left}px`
    popelement.style.top = `${top}px`
    if (right != null) popelement.style.right = `${right}px`
    popelement.addEventListener("mouseover", onIframeHover)
    popelement.addEventListener("mouseout", onIframeOut)
    document.body.appendChild(popelement)
}

function closeIframe() {
    if (popelement != null) {
        document.body.removeChild(popelement)
        popelement = null
        mouseOnFrame = false
        hoverTriggerElement = null
    }
}
let mouseOnFrame = false
let lastDocOut = -1
function onDoclinkOut(node) {
    hoverTriggerElement = null
    trigger.id += 1
    let id = trigger.id
    lastDocOut = setTimeout(() => {
        if (!mouseOnFrame && id === trigger.id) closeIframe()
        lastDocOut = -1
    }, 1000);
}

function onIframeHover() {
    mouseOnFrame = true
}

function onIframeOut() {
    mouseOnFrame = false
    closeIframe()
}

function isRootWindow() {
    return window.self == window.parent
}

window.addEventListener("message", handleIframeMessage)

function handleIframeMessage(e) {
    let node = e.data.node
    let url = e.data.url
    let p = e.data.p
    if (isRootWindow()) {
        createIframe(null, url, p)
    } else {
        window.parent.postMessage({ node: node, url, p }, "*")
    }
}