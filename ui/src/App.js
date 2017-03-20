import React, { Component } from 'react'
import logo from './logo.svg'
import './App.css'
import fetch from 'isomorphic-fetch'
import Payload from './Payload'

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      data: [],
      err: '',
    }
  }

  getData() {
    fetch('http://localhost:9090/api/docs')
      .then(res => res.json())
      .then(json => this.setState({data: json}))
      .catch(err => this.setState({err}))
  }

  componentDidMount() {
    this.getData()
  }

  render() {
    const { data, err } = this.state
    console.log(data)
    return (
      <div className="App">
        <div className="App-header">
          <h2>Docs.</h2>
        </div>
        <p className="App-intro">
          {(data ? data.map(item =>
            <div>
              <div className="code-meta">
                <p className="path">{item.Path}</p>
                <p>{item.Method} - {item.Handler}</p>
              </div>
              {(item.Params ? <Payload item={item} /> : false)}
            </div>
          ) : false)}
        </p>
      </div>
    );
  }
}

export default App
