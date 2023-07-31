
const imageUploadInput = document.getElementById("imageUpload");
const progressBar = document.getElementById("image-upload-progress");

imageUploadInput.onchange = () => {
    progressBar.style.width = "0%";
    setTimeout(() => {
        const xhr = new XMLHttpRequest();
        progressBar.classList.add("progress-bar-striped");
        progressBar.classList.remove("bg-success");
        xhr.addEventListener("progress", e => {
            progressBar.style.width = (e.loaded / e.total * 100) + "%";
        });
        xhr.addEventListener("load", () => {
            setTimeout(() => {
                progressBar.classList.remove("progress-bar-striped");
                progressBar.classList.add("bg-success");
            }, 500);
        });
        xhr.open("POST", "/imageupload");
        const formData = new FormData();
        formData.append("file", imageUploadInput.files[0]);
        xhr.send(formData);
    }, 500);
};
