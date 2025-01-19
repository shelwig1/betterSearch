// import './style.css';
// import './app.css';

import logo from './assets/images/logo-universal.png';
import {ChooseDirectory, GetDirectoryMap, Greet} from '../wailsjs/go/main/App';

const chooseDirButton = document.getElementById("choose-dir-button")

let currentDirectoryMap

chooseDirButton.addEventListener("click", async () => {
    console.log("Directory button clicked")

    try {
        const result = await ChooseDirectory()

        console.log("Full result: ", result)
        currentDirectoryMap = GetDirectoryMap(result)
        
    } catch (err) {
        console.log("Catch triggered: ", err)
    }

})

window.greet = function () {

    console.log(GoHello())

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
