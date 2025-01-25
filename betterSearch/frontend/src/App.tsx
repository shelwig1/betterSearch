import React, { useState, useEffect } from 'react';
import ReactDOM from 'react-dom';

interface SearchResultProp {
    target : string;
    file : string;
}
const App = () => {
    return (
        <div>
            <h1>Hello from React in Wails!</h1>
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
        <div className='searchResultCard'>
            <p>{target}</p>
            <p>{file}</p>
        </div>
    )
}
//ReactDOM.render(<App />, document.getElementById('app'))



export default App;