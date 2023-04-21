
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
