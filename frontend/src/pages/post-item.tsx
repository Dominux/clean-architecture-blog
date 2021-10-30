import React from 'react'
import { Card, CardContent, Typography } from '@mui/material'

import { PostData } from './types'

const Post: React.FC<PostData> = (props: PostData) => {
  return (
    <Card>
      <CardContent>
        <Typography variant="h5">{props.author}</Typography>
        <Typography variant="h4">{props.title}</Typography>
        <Typography>{props.text}</Typography>
      </CardContent>
    </Card>
  )
}

export default Post
