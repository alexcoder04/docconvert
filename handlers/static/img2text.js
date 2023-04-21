
function copyOriginal() {
    let text = document.getElementById("myText").value;
    if (text == "") {
        return;
    }
    navigator.clipboard.writeText(text).then(() => {
        document.querySelector("#btn-copy-original").innerHTML = "<i class='bi-clipboard'> Copied!</i>"
        setTimeout(() => {
            document.querySelector("#btn-copy-original").innerHTML = "<i class='bi-clipboard'></i> Copy original text"
        }, 3000);
    });
}

function copyChatGPT() {
    let text = document.getElementById("myText").value;
    if (text == "") {
        return;
    }
    text = `Please summarize the following text:\n${text}`;
    navigator.clipboard.writeText(text).then(() => {
        document.querySelector("#btn-copy-chatgpt").innerHTML = "<i class='bi-clipboard'> Copied!</i>"
        setTimeout(() => {
            document.querySelector("#btn-copy-chatgpt").innerHTML = "<i class='bi-clipboard'></i> Copy ChatGPT summarize prompt"
        }, 3000);
    });
}

function downloadTxt() {
    let text = document.getElementById("myText").value;
    let element = document.createElement("a");
    element.setAttribute("href", "data:text/plain;charset=utf-8," + encodeURIComponent(text));
    element.setAttribute("download", "yourtext.txt");
    element.style.display = "none";
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
}

document.addEventListener("DOMContentLoaded", () => {
    const tabs = document.querySelectorAll('.nav-tabs > li > button');

    tabs.forEach(tab => {
        tab.addEventListener("click", () => {
            localStorage.setItem("activeTab", tab.getAttribute("data-bs-target"));
        });
    });

    const activeTab = localStorage.getItem("activeTab");

    if (activeTab) {
        const tab = document.querySelector(`.nav-tabs button[data-bs-target="${activeTab}"]`);
        tab.click();
    }
});
