let started = false;

function getTextArea() {
    return input = document.getElementsByClassName("ace_text-input")[0];
}

function init() {
    textArea = getTextArea();
    let iframe = document.createElement("iframe");
    iframe.style = "position: absolute;" +
        "left: 50px;" +
        "bottom: 0px;" +
        "background: rgba(0, 0, 0, 0.1);" +
        "height: 186px;" +
        "width: 506px;" +
        "z-index: 100;";

    // Note: uncomment for development
    // iframe.src = "https://localhost:8080/";
    iframe.src = "https://latex.davidw.tech/";
    textArea.parentElement.appendChild(iframe);

    window.addEventListener("message", function(e) {
        // Find the ace editor
        let x = document.getElementsByClassName("ace-editor-body");
        let ace = window.wrappedJSObject.ace;
        let editor = ace.edit(x[0]);

        // Insert the output where the cursor is.
        editor.session.insert(
            editor.getCursorPosition(),
            e.data.output,
        );
    });
    started = true;
}

function reinit() {
    if (!getTextArea()) {
        setTimeout(reinit, 100);
    } else {
        // Delay for a second after finding the initial window.
        setTimeout(init, 1000);
    }
}
reinit();
