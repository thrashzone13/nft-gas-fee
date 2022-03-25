import React, { useEffect, useState } from "react";
import { Accordion, Button, ButtonGroup, ListGroup } from "react-bootstrap";
import _ from "lodash";

function App() {
  const [data, fetchData] = useState({})
  const [type, setType] = useState('average')

  const getData = () => {
    var url = !process.env.NODE_ENV || process.env.NODE_ENV === 'production' ? 'docker.internal.host/price' : 'http://localhost:8080/price';
    fetch(url)
      .then(res => res.json())
      .then(res => fetchData(res))
  }

  useEffect(() => {
    getData()
  }, [])

  const {
    contracts = {},
    eth = 0,
    gwei = {}
  } = data;

  return (
    <>
      <ButtonGroup size="lg" className="mb-2">
        <Button onClick={()=>setType('low')}>Low</Button>
        <Button onClick={()=>setType('average')}>Average</Button>
        <Button onClick={()=>setType('fast')}>Fast</Button>
      </ButtonGroup>
      <Accordion>
        {Object.keys(contracts).map((c, i) => <Accordion.Item eventKey={i}>
          <Accordion.Header>{_.upperCase(c)}</Accordion.Header>
          <Accordion.Body>
            <ListGroup variant="flush">
              {Object.keys(contracts[c]).map((v, i) =>
                <ListGroup.Item
                  key={i}
                  as="li"
                  className="d-flex justify-content-between align-items-start"
                >
                  <p>{_.upperCase(v)}</p>
                  <p>{((gwei[type] * contracts[c][v] * 0.000000001) / eth).toFixed()} $</p>
                </ListGroup.Item>
              )}
            </ListGroup>
          </Accordion.Body>
        </Accordion.Item>
        )}
      </Accordion>
    </>

  );
}

export default App;