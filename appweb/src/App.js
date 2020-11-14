import React, { Component } from 'react';
import './App.css';

import ChartistGraph from 'react-chartist'

class App extends Component {

  url_api = 'http://localhost:3000/'

  constructor() {
    super()
    this.state = {
      topCases: [
        {
          _id: "Guatemala",
          number: 3000
        },
        {
          _id: "Peten",
          number: 2000
        },
        {
          _id: "El Progreso",
          number: 1000
        }
      ],
      allCases: {
        labels: [],
        series: []
      },
      rangeCases: {
        labels: [],
        series: [
          []
        ]
      },
      lastCase: {
        name: "Cindy Lu",
        location: 'Villa Quien',
        age: 32,
        infectedtype: "community",
        state: 'asymptomatic'
      }
    }

    fetch(this.url_api + 'cases/getCasesPerLocation').then(res => {
      return res.json()
    }).then(data => {
      if (data) {
        var a = []
        for (let index = 0; index < data.data.length; index++) {
          if (index < 3) {
            a.push(data.data[index]);
          }
          this.state.allCases.labels.push(data.data[index]._id)
          this.state.allCases.series.push({ value: data.data[index].number, name: data.data[index]._id })
        }

        this.setState({
          topCases: a,
          allCases: this.state.allCases
        })
      }
    })

    fetch(this.url_api + 'cases/getCasesPerRange').then(res => {
      return res.json()
    }).then(data => {
      if (data) {

        var a = [0]
        var l = ["1-10"]
        let i = 1;
        let aux = i * 10;
        for (let index = 0; index < data.data.length; index++) {
          if (data.data[index]._id < aux) {
            a[0] = a[0] + data.data[index].number;
          } else {
            while (data.data[index]._id > aux) {
              i++
              aux = i * 10;
              a.push(0);
              l.push((aux - 9) + "-" + aux);
            }
            //i++
            //a.push(0);
            //l.push((aux-9)+"-"+aux);
            a[i - 1] = a[i - 1] + data.data[index].number;
          }

          //this.state.allCases.labels.push(data.data[index]._id)
          //this.state.allCases.series.push({ value: data.data[index].number, name: data.data[index]._id })
        }

        this.state.rangeCases.series[0] = a;
        this.state.rangeCases.labels = l;

        this.setState({
          rangeCases: this.state.rangeCases
        })
      }
    })

    fetch(this.url_api + 'cases/getLastMongo').then(res => {
      return res.json()
    }).then(data => {
      if (data) {

        this.state.lastCase = data.data[0];

        this.setState({
          lastCase: this.state.lastCase
        })
        console.log(data.data[0])
      }
    })
  }

  render() {

    return (
      <div className="App">
        <nav className="navbar navbar-dark bg-dark">
          <span className="navbar-brand">CoronaInfo</span>
        </nav>
        <div className="container">
          <div className="text-center">
            <h1 className="display-3">
              Coronavirus
          </h1>
          </div>
          <div className="dropdown-divider"></div>
          <div className="row mt-3">
            <div className="col-12 col-md-5">
              <div className="card text-center">
                <div className="card-header">Top 3: Departamentos con Coronavirus</div>
                <ul className="list-group list-group-flush">
                  {
                    this.state.topCases.map((dept, i) => {
                      return (
                        <li className="list-group-item" key={i}>{dept._id}: {dept.number} Caso(s)</li>
                      )
                    })
                  }
                </ul>
              </div>
            </div>
            <div className="col-12 col-md-1">

            </div>
            <div className="col-12 col-md-5">
              <div className="border border-info">
                <div className="card-header">Departamentos Afectados</div>
                <ChartistGraph data={this.state.allCases} options={
                  {
                    width: "300px",
                    height: "300px",
                    donut: false,
                  }
                } type={'Pie'} />
              </div>
            </div>
          </div>
          <div className="dropdown-divider mt-3"></div>
          <div className="row mt-3">
            <div className="col-12 col-md-5">
              <div className="card text-center">
                <div className="card-header">Ultimo Caso Agregado</div>
                <ul className="list-group list-group-flush text-left">
                  <li className="list-group-item"><strong>Nombre:</strong> {this.state.lastCase.name}</li>
                  <li className="list-group-item"><strong>Lugar:</strong> {this.state.lastCase.location}</li>
                  <li className="list-group-item"><strong>Edad:</strong> {this.state.lastCase.age}</li>
                  <li className="list-group-item"><strong>Tipo de Infeccion:</strong> {this.state.lastCase.infectedtype}</li>
                  <li className="list-group-item"><strong>Estado:</strong> {this.state.lastCase.state}</li>
                </ul>
              </div>
            </div>
            <div className="col-12 col-md-1">

            </div>
            <div className="col-12 col-md-5">
              <div className="border border-info mb-4">
                <div className="card-header">Edades Afectadas</div>
                <ChartistGraph data={this.state.rangeCases} options={
                  {
                    reverseData: true,
                    horizontalBars: true,
                    width: "100%",
                    height: "300px",
                  }
                } type={'Bar'} />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default App;
