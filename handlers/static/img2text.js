
const copyOriginalBtn = document.getElementById("btn-copy-original");
const copyChatGptBtn = document.getElementById("btn-copy-chatgpt");

addCopyFunction(copyOriginalBtn, () => {
    return document.getElementById("myText").value;
});

addCopyFunction(copyChatGptBtn, () => {
    const text = document.getElementById("myText").value;
    if (text.trim() == "") {
        return "";
    }
    return `Please summarize the following text:\n${text}`;
});

function downloadTxt() {
    const text = document.getElementById("myText").value;
    const element = document.createElement("a");
    element.setAttribute("href", "data:text/plain;charset=utf-8," + encodeURIComponent(text));
    element.setAttribute("download", "yourtext.txt");
    element.style.display = "none";
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
}
