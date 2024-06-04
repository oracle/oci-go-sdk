'use strict';

const e = React.createElement;

class LikeButton extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            liked: false ,
            masterItems: "",
            localItems:""
        };
    }

    componentDidMount(){
        fetch(`https://objectstorage.us-phoenix-1.oraclecloud.com/p/5mzix52OhxjnITDKe5bdefXAwxOLpEUEta1czeh_aK4/n/dex-us-phoenix-1/b/codecov_baseline/o/codecov_master.json`)
            .then(result=>result.json())
            .then(items=>this.setState({masterItems: items}))
        fetch(`https://objectstorage.us-phoenix-1.oraclecloud.com/p/y4R2h_AwvDBBG0avDuy8ZilmHWQU8MrGP2GYadWP91Y/n/dex-us-phoenix-1/b/codecov_baseline/o/codecov_preview.json`)
            .then(result=>result.json())
            .then(items=>this.setState({previewItems: items}))
        fetch('./master_local_codecov.json')
            .then(result=>result.json())
            .then(items=>this.setState({localItems: items}))
        // var localData = require('./local_codecov.xml')

    }

    render() {
        if (this.state.liked) {
            var itemsArray = this.getItemArray(this.state.masterItems);
            var previewArray = this.getItemArray(this.state.previewItems);
            var localArray = this.getItemArray(this.state.localItems);

            itemsArray.sort((a, b) => (a.name > b.name) ? 1 : -1);
            previewArray.sort((a, b) => (a.name < b.name) ? 1 : -1);
            localArray.sort((a, b) => (a.name < b.name) ? 1 : -1);

            return (
                <div>
                <h3>Master branch</h3>
                {this.getTable(itemsArray)}
                <h3>Preview branch</h3>
                {this.getTable(previewArray)}
                <h3>Local</h3>
                    {this.getTable(localArray)}
                </div>
            );
        }
        return e(
            'button',
            { onClick: () => this.setState({ liked: true }) },
            'Generate'
        );
    }

    // given an array of items stored in the state, return an array that will be used to display the table
    getItemArray(items) {

        var itemsArray = [];
        for (var key in items) {
            for(var i=0; i < items[key].length; i++) {
                var innerItem = items[key][i];
                var item = {};
                item.name = key;
                item.tag = innerItem.tag;
                item.className = innerItem.class;
                var entry = innerItem.data;
                item.instructionCoverage = Math.round(entry.coveredInstructions * 100 / (entry.coveredInstructions + entry.missedInstructions));
                item.branchCoverage = Math.round(entry.coveredBranches * 100 / (entry.missedBranches + entry.coveredBranches));
                item.lineCoverage = Math.round(entry.coveredLines * 100 / (entry.coveredLines + entry.missedLines));
                var d = new Date(Date.parse(entry.timestamp));
                item.timestamp = d.toLocaleString();
                item.testsPassedPercentage = entry.testsPassedPercentage;
                itemsArray.push(item);
            }
        }

        return itemsArray;
    }

    // given an item array, return table
    getTable(itemsArray) {
        const header = ["service", "tag", "class", "Instruction coverage", "Branch coverage", "Line coverage", "Timestamp", "Tests Passed (%)"];

        return (
            <table>
                <thead>
                <tr>{header.map((h, i) => <th key={i}>{h}</th>)}</tr>
                </thead>
                <tbody>{itemsArray.map(function(item, key) {

                    return (
                        <tr key = {key}>
                            <td>{item.name}</td>
                            <td>{item.tag}</td>
                            <td>{item.className}</td>
                            <td>{item.instructionCoverage}</td>
                            <td>{item.branchCoverage}</td>
                            <td>{item.lineCoverage}</td>
                            <td>{item.timestamp}</td>
                            <td>{item.testsPassedPercentage}</td>
                        </tr>
                    )

                })}
                </tbody>
            </table>
        )
    }
}

const domContainer = document.querySelector('#generate_button_container');
ReactDOM.render(e(LikeButton), domContainer);
