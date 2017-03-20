import React from 'react'

export const Payload = (props) =>
  <code className="param-container">
    <span dangerouslySetInnerHTML={{__html: '{' }}></span>
      <div className="params">{props.item.Params.map(param => <p>{`${param.Tag}: ${param.Type}`}</p>)}</div>
    <span dangerouslySetInnerHTML={{__html: '}' }}></span>
  </code>

export default Payload
