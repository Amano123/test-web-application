// import React from 'react';
// import ReactDOM from 'react-dom';
// import './index.css';
// import App from './App';
// import reportWebVitals from './reportWebVitals';

// ReactDOM.render(
//   <React.StrictMode>
//     <App />
//   </React.StrictMode>,
//   document.getElementById('root')
// );

// // If you want to start measuring performance in your app, pass a function
// // to log results (for example: reportWebVitals(console.log))
// // or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
// reportWebVitals();
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
    fetch("http://localhost:4000") //api
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
            {Object.keys(items).map(key => (
              <li key={key}>{key} - {items[key]}</li>
            ))}
          </ul>
        </div>
      );
    }
  }
}

export default Exchange;

render(<Exchange />, document.getElementById('root'));
