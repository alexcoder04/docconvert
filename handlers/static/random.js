
// NUMBERS
const generateNumberButton = document.getElementById("generateNumber");
const minNumberInput = document.getElementById("minNumber");
const maxNumberInput = document.getElementById("maxNumber");
const generatedNumberInput = document.getElementById("generatedNumber");
const copyButton = document.getElementById("copyButton");

generateNumberButton.addEventListener("click", function() {
    const minNumber = parseInt(minNumberInput.value);
    const maxNumber = parseInt(maxNumberInput.value);
    const randomNumber = Math.floor(Math.random() * (maxNumber - minNumber + 1)) + minNumber;
    generatedNumberInput.value = randomNumber;

    generatedNumberInput.classList.add("is-valid");
    setTimeout(() => {
        generatedNumberInput.classList.remove("is-valid");
    }, 1000);
});

addCopyFunction(copyButton, () => {
    return generatedNumberInput.value;
});


// NAMES
const newNameInput = document.getElementById("new-name");
const addNameButton = document.getElementById("add-name");
const nameList = document.getElementById("name-list");
const randomNameDisplay = document.getElementById("pickedName");
const pickRandomButton = document.getElementById("pick-random");
const copyNameButton = document.getElementById("copyNameButton");

newNameInput.addEventListener("keydown", (event) => {
    if (event.key == "Enter") {
        addNameButton.click();
        return;
    }
});

function addNameToDOM(name) {
    const newNameItem = document.createElement("li");
    newNameItem.classList.add("list-group-item", "d-flex", "justify-content-between", "align-items-center");
    const newNameText = document.createElement("span");
    newNameText.textContent = name;
    const removeNameButton = document.createElement("button");
    removeNameButton.classList.add("btn", "btn-danger", "btn-sm", "remove-name-button");
    removeNameButton.innerHTML = "<i class='bi-trash'></i> Remove";
    newNameItem.appendChild(newNameText);
    newNameItem.appendChild(removeNameButton);
    nameList.appendChild(newNameItem);
    newNameInput.value = "";
}

addNameButton.addEventListener("click", () => {
    const newName = newNameInput.value.trim();
    if (newName) {
        addNameToDOM(newName);
        lsArrAdd("names", newName);
    }
});

nameList.addEventListener("click", (event) => {
    if (event.target.tagName === "BUTTON") {
        lsArrRemove("names", event.target.parentElement.querySelector("span").textContent);
        event.target.parentElement.remove();
    }
});

pickRandomButton.addEventListener("click", () => {
    const nameItems = nameList.querySelectorAll("li");
    const randomIndex = Math.floor(Math.random() * nameItems.length);
    randomNameDisplay.value = nameItems[randomIndex].querySelector("span").textContent;
});


addCopyFunction(copyNameButton, () => {
    return randomNameDisplay.value;
});

lsArrGetFull("names").forEach(name => {
    addNameToDOM(name);
});
