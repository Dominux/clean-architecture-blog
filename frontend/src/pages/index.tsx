import React, { useState, useEffect } from 'react'
import axios from 'axios'

import { PostData } from './types'
import PostItem from './post-item'

const getPosts = async (): Promise<Array<PostData>> => {
  const response = await axios.get(`/api/posts/`)
  return response.data.message.filter((e: PostData | null) => e)
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
