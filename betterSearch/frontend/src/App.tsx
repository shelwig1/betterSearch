import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';

import './style.css';
import './app.css';

interface SearchResultProp {
    target : string;
    file : string;
}
const App = () => {
    return (
        <div>
            <div id="dir-ops-container">
                <div id="current-dir">Current directory: </div>
                <div id="search-ready-status">Not ready to search</div>
                <div className="throbber" id="dir-map-throbber"></div>        
            </div>

        <button id="choose-dir-button">Choose Directory</button>
        <TestButton />
        <SearchResultContainer />
        </div>
    )
}
// Want to make it so as we spawn these cocksuckers, they get routed into here
// Need to hold the function call for the search, then pass it on to the kids
const SearchResultContainer = () => {
    const [results, setResults] = useState<SearchResultProp[]>([])

    const handleNewResult = (newResult: SearchResultProp) => {
        setResults((prevResults) => [...prevResults, newResult])
    }
    // This is where we make our money - import the search function here, listen as we stream in more and more shit, ezpz
/*      useEffect(() => {
        const searchData = async () => {
            const newResults = await fetchSearchResults()
            handleNewResult(newResults)
        }
    }) */
    useEffect(() => {
        const fetchData = () => {
            const newResults = mockSearchData()
            setResults(newResults)
        }
        fetchData()
    }, [])


    return (
        <div id='searchResultsContainer'>
            {results.map((result, index) => (
                <SearchResultCard key={index} target={result.target} file={result.file} />
            ))}
        </div>
    )
}
const SearchResultCard: React.FC<SearchResultProp> = ({target, file}) => {
    return (
        <button className='searchResultCard' onClick={() => console.log("result testicles")}>
            <p>{target}</p>
            <p>{file}</p>
        </button>
    )
}
//ReactDOM.render(<App />, document.getElementById('app'))

const TestButton: React.FC = () => {
    return (
        <button onClick={() => console.log("Testicles")}>
            Testicles Button
        </button>
    )
}

const mockSearchData = () => {
    // Simulating some mock data
    return [
      { target: 'Target 1', file: 'File 1' },
      { target: 'Target 2', file: 'File 2' },
      { target: 'Target 3', file: 'File 3' },
    ];
  };


export default App;