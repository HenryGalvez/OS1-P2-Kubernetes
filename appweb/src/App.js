import React, { Component } from 'react';
import './App.css';

class App extends Component {



  constructor() {
    super()
    this.state = {
      topCases: [
        {
          name: "Guatemala",
          number: 3000
        },
        {
          name: "Peten",
          number: 2000
        },
        {
          name: "El Progreso",
          number: 1000
        }
      ],
      lastCase: {
        name: "Cindy Lu",
        location: 'Villa Quien',
        age: 32,
        infectedtype: "community",
        state: 'asymptomatic'
      }
    }
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
                        <li className="list-group-item" key={i}>{dept.name}: {dept.number} Caso(s)</li>
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
                <img src="https://d2mvzyuse3lwjc.cloudfront.net/doc/en/UserGuide/images/2D_B_and_W_Pie_Chart/2D_B_W_Pie_Chart_1.png?v=83139" alt="" className="img-fluid" />
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
              <div className="border border-info">
                <img src="https://apexcharts.com/wp-content/uploads/2018/05/basic-bar-chart.svg" alt="" className="img-fluid" />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default App;
