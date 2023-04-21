
var generateNumberButton = document.getElementById("generateNumber");
var minNumberInput = document.getElementById("minNumber");
var maxNumberInput = document.getElementById("maxNumber");
var generatedNumberInput = document.getElementById("generatedNumber");
var copyButton = document.getElementById("copyButton");

generateNumberButton.addEventListener("click", function() {
    var minNumber = parseInt(minNumberInput.value);
    var maxNumber = parseInt(maxNumberInput.value);
    var randomNumber = Math.floor(Math.random() * (maxNumber - minNumber + 1)) + minNumber;
    generatedNumberInput.value = randomNumber;

    generatedNumberInput.classList.add("is-valid");
    setTimeout(() => {
        generatedNumberInput.classList.remove("is-valid");
    }, 1000);
});

copyButton.addEventListener("click", function() {
    navigator.clipboard.writeText(generatedNumberInput.value).then(() => {
        copyButton.innerHTML = "<i class='bi-clipboard'> Copied!</i>"
        setTimeout(() => {
            copyButton.innerHTML = "<i class='bi-clipboard'></i> Copy"
        }, 3000);
    });
});
