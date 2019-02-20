
let canvas = null;
let canvasCtx = null;
let canvasHeight = 0;
let canvasWidth = 0;
let drawing = false;

function submitCanvasInput() {
    console.log("the data: ", canvas.toDataURL("image/jpeg"))
}

function clearCanvasInput() {
    canvasCtx.clearRect(0, 0, canvasWidth, canvasHeight)
}

function initCanvas() {
    canvas = document.getElementById("input-canvas");
    canvasCtx = canvas.getContext('2d');
    canvas.height = canvas.clientHeight;
    canvas.width = canvas.clientWidth;
    canvasHeight = canvas.height;
    canvasWidth = canvas.width;

    canvas.addEventListener("mousedown", startDraw);
    canvas.addEventListener("mousemove", paintAtMouse);
    canvas.addEventListener("mouseup", endDraw);
    canvas.addEventListener("mouseout", endDraw);

    // TODO: Also need a resize I think
}

function startDraw(e) {
    drawing = true;
    paintAtMouse(e)
}

let lastX, lastY = [null, null];

function paintAtMouse(e) {
    if (!drawing) {
        return
    }

    canvasCtx.strokeStyle = "#000";
    let x = parseInt(e.clientX - canvas.offsetLeft);
    let y = parseInt(e.clientY - canvas.offsetTop);

    if (lastX == null) {
        // Just draw a box
        canvasCtx.fillRect(x - 1, y - 1, 2, 2);
    } else {
        lineFromLast(e);
    }
    [lastX, lastY] = [x, y];
}

function endDraw(e) {
    if (lastX != null) {
        // We need to finish the current line
        lineFromLast(e);
    }

    [lastX, lastY] = [null, null];
    drawing = false;
}

function lineFromLast(e) {
    canvasCtx.strokeStyle = "#000";
    let x = parseInt(e.clientX - canvas.offsetLeft);
    let y = parseInt(e.clientY - canvas.offsetTop);

    canvasCtx.beginPath();
    canvasCtx.lineWidth = 2;
    canvasCtx.moveTo(lastX, lastY);
    canvasCtx.lineTo(x, y);
    canvasCtx.stroke();
}

initCanvas();
