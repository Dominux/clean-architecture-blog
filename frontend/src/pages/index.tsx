import React, { useState, useEffect } from 'react'
import axios from 'axios'

import { PostData } from './types'
import PostItem from './post-item'

const getPosts = async (): Promise<Array<PostData>> => {
  return axios.get('localhost:8080/api/posts')
}

const Home: React.FC = () => {
  const [posts, setPosts] = useState<PostData[]>([])

  // Fetching data from server once the page mounted
  useEffect(() => {
    ;(async () => setPosts(await getPosts()))()
  }, [])

  return (
    <React.Fragment>
      {posts.map((post) => (
        <PostItem
          author={post.author}
          title={post.title}
          text={post.text}
        ></PostItem>
      ))}
    </React.Fragment>
  )
}

export default Home
