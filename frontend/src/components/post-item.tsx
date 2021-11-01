import React from 'react'
import { Card, CardContent, Typography} from '@mui/material'
import {Link} from 'react-router-dom'

import { PostData } from '../pages/types'

const PostItem: React.FC<PostData> = (props) => {
  return (
    <Card sx={{ my: 3 }}>
      <CardContent>
        <Link to={`/posts/${props.id}`}>
          <Typography variant="h5">{props.author}</Typography>
        </Link>
        <Typography variant="h4">{props.title}</Typography>
        <Typography>{props.text}</Typography>
      </CardContent>
    </Card>
  )
}

export default PostItem
