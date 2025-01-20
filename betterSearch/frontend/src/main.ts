import './style.css';
import './app.css';

// import logo from './assets/images/logo-universal.png';
import {ChooseDirectory, GetDirectoryMap, Greet} from '../wailsjs/go/main/App';

const chooseDirButton = document.getElementById("choose-dir-button") as HTMLButtonElement
const currentDirDisplay = document.getElementById("current-dir") as HTMLDivElement
const searchReadyStatus = document.getElementById("search-ready-status") as HTMLDivElement
const throbber = document.getElementById("throbber") as HTMLElement

let currentDirectory
let currentDirectoryMap


chooseDirButton.addEventListener("click", async () => {
    console.log("Directory button clicked")

    try {
        const result = await ChooseDirectory()

        console.log("Full result: ", result)
        setCurrentDirectory(result)
        
    } catch(err) {
        console.log("Catch triggered: ", err)
    }

})

async function setCurrentDirectory(dir : string) {
    searchReadyStatus.innerHTML = "Not ready to search"

    throbber.style.display = "block"


    currentDirectory = dir

    currentDirDisplay.innerHTML = "Current directory: " + dir

    currentDirectoryMap = await GetDirectoryMap(dir)

    throbber.style.display = "none"

    searchReadyStatus.innerHTML = "Ready to search!"

    // Update the text to reflect what we give a fuck about
}



/* window.greet = function () {


    // Get name
    let name = nameElement.value;

    // Check if the input is empty
    if (name === "") return;

    // Call App.Greet(name)
    try {
        Greet(name)
            .then((result) => {
                // Update result with data back from App.Greet()
                resultElement.innerText = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};
 */