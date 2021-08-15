const canvas = document.getElementById('canvas')

const ctx = canvas.getContext('2d')

const Rect = {
    ctx: ctx,
    x: canvas.width/2,
    y: canvas.height/2,
    w: 50,
    h: 50,
    update(x,y) {
        this.x = x
        this.y = y
    },
    draw() {
        ctx.fillRect(this.x, this.y, this.w, this.h)
    }
}

const clear = () => {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
}

const drawBorder = () => {
    ctx.strokeRect(0, 0, canvas.width, canvas.height)
}

const updateRect = (x, y) => {
    clear()
    drawBorder()
    Rect.update(x,y)
    Rect.draw()
}

const init = () => {
    drawBorder()
    Rect.draw()
}

const enableMoving = () => {
    let moving = false
    let offsetX, offsetY

    document.addEventListener("mousedown", (e) => {
        moving = true
        offsetX = e.offsetX - Rect.x
        offsetY = e.offsetY - Rect.y
    })

    document.addEventListener("mousemove", (e) => {
        if (moving) {
            updateRect(e.offsetX - offsetX, e.offsetY - offsetY)
        }
    })

    document.addEventListener("mouseup", () => moving = false)
}

init()
enableMoving()