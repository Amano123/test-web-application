import React, { Component } from 'react';
import { render } from 'react-dom';

class Exchange extends Component {
  constructor(props) {
    super(props);
    this.state = { //state初期化
      isLoaded: false,
      items: []
    };
  }
  componentDidMount() { //render直後に行いたい処理を書くところ
    fetch("http://localhost:1234") //api
      .then(res => res.json()) 
      .then(json => {
        console.log(json.rates);
        this.setState({
          isLoaded: true,
          items: json
        });
      });
  }

  render() {
    var { items, isLoaded } = this.state;
console.log(items);
    if (!isLoaded) {
      return <div>...Loading</div>;
    } else {
      return (
        <div>
          <ul>
            {Object.keys(items[1]).map(key => (
              <li key={key}>{key} - {items[1][key]}</li>
            ))}
            {Object.keys(items[2]).map(key => (
              <li key={key}>{key} - {items[2][key]}</li>
            ))}
          </ul>
        </div>
      );
    }
  }
}

export default Exchange;

render(<Exchange />, document.getElementById('root'));