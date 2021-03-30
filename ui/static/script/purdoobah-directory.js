// search bar, attach type input listener
const inputSearch = document.getElementById("inputSearch");
inputSearch.addEventListener("input", search);

// the archives-incomplete gif
// set it to invisible initially
const archivesIncomplete = document.getElementsByClassName("archives-incomplete")[0];
archivesIncomplete.style.display = "none";

// `search` shows/hides Purdoobah Cards based on the search term
function search() {
    const search = this.value.toLowerCase();

    const purdoobahCards = document.getElementsByClassName("purdoobah-card");

    let cardsHidden = 0;
    for (let i = 0; i < purdoobahCards.length; i++) {
        const purdoobahCard = purdoobahCards.item(i);

        const name = purdoobahCard.dataset.name;
        const emoji = purdoobahCard.dataset.emoji;
        const birthCertificateName = JSON.parse(purdoobahCard.dataset.birthCertificateName);
        const yearsMarched = purdoobahCard.dataset.yearsMarched;

        if (containsSearch(search, name, emoji, birthCertificateName, yearsMarched)) {
            purdoobahCard.style.display = "block";
        } else {
            purdoobahCard.style.display = "none";
            cardsHidden++;
        }
    }

    // if every purdoobah card is hidden, show archives-incomplete gif
    if (cardsHidden === purdoobahCards.length) {
        archivesIncomplete.style.display = "block";
    } else {
        archivesIncomplete.style.display = "none";
    }
}

// `containsSearch` returns true if search term can be found in any of the provided values
function containsSearch(search, name, emoji, birthCertificateName, yearsMarched) {
    // name
    if (name.toLowerCase().includes(search)) {
        return true;
    }

    // emoji
    if (emoji.includes(search)) {
        return true;
    }

    // first name
    if (birthCertificateName.first.toLowerCase().includes(search)) {
        return true;
    }

    // middle name
    if (birthCertificateName.hasOwnProperty("middle") &&
        birthCertificateName.middle.toLowerCase().includes(search)) {
        return true;
    }

    // last name
    if (birthCertificateName.last.toLowerCase().includes(search)) {
        return true;
    }

    // years marched
    if (yearsMarched.includes(search)) {
        return true;
    }

    return false;
}
