import React, { useState, useEffect } from 'react'

const ApiFetch = () => {

  const [posts, setPosts] = useState([])

  useEffect(() => {
    fetch('http://localhost:1234', { method: 'GET' })
      .then(res => {
        posts(res)
      })
  }, [])

  return (
    <div>
      <ul>
        {setPosts}
      </ul>
    </div>
  )
}

export default ApiFetch