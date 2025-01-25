import React from 'react';
import ReactDOM from 'react-dom/client'; 

import './style.css';
import './app.css';

import App from './App.js'
// import logo from './assets/images/logo-universal.png';
import {ChooseDirectory, GetDirectoryMap, Greet} from '../wailsjs/go/main/App.js';

const chooseDirButton = document.getElementById("choose-dir-button") as HTMLButtonElement
const currentDirDisplay = document.getElementById("current-dir") as HTMLDivElement
const searchReadyStatus = document.getElementById("search-ready-status") as HTMLDivElement
const dirMapThrobber= document.getElementById("dir-map-throbber") as HTMLElement

let currentDirectory
let currentDirectoryMap

console.log("Test of this doodad")
const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

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

// Current implementation is weird - may turn this into a React project and make everything "cleaner"
async function setCurrentDirectory(dir : string) {
    searchReadyStatus.innerHTML = "Not ready to search"

    dirMapThrobber.style.display = "block"

    currentDirectory = dir

    currentDirDisplay.innerHTML = "Current directory: " + dir

    // Need to figure out a way to handle errors on this one too
    // Figure out if this breaks anything
    try {
        currentDirectoryMap = await GetDirectoryMap(dir)
    } catch (err) {
        searchReadyStatus.innerHTML = "Unable to access directory" 
        return
    }

    dirMapThrobber.style.display = "none"

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