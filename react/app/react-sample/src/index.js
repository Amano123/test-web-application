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
    fetch("http://localhost:3000") //api
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
      var obj = JSON.stringify(items);
      return (
        <div style={{
          height: "calc(100vh - 2 * 80px)",
        }}>
          this is sample.
        <p>
          {obj}
        </p>
        </div>
      );
    }
  }
}

export default Exchange;

render(<Exchange />, document.getElementById('root'));