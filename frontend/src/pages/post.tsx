import React from 'react'
import { useParams } from 'react-router-dom'

const Post: React.FC = () => {
  let { id } = useParams<{id: string}>()
  return <React.Fragment>{ id }</React.Fragment>
}

export default Post
