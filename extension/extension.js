
function init() {
    let input = document.getElementsByClassName("ace_text-input");
    if (!input) {
        console.error("input not found");
        return;
    }

    textArea = input[0];
    let iframe = document.createElement("iframe");
    iframe.style = "position: absolute;" +
        "left: 50px;" +
        "bottom: 0px;" +
        "background: rgba(0, 0, 0, 0.1);" +
        "height: 179px;" +
        "width: 306px;";
    iframe.src = "https://localhost:8080/";
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
    })
}

init();
