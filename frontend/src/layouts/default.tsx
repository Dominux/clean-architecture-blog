import {
  AppBar,
  Container,
  IconButton,
  Toolbar,
  Typography,
} from '@mui/material'
import { Box } from '@mui/system'
import { Menu } from '@mui/icons-material'
import React from 'react'

const DefaultLayout: React.FC = (props) => {
  return (
    <React.Fragment>
      {/* Header */}
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar variant="dense">
            <Typography variant="h5" color="inherit" component="div">
              Blog
            </Typography>
            <IconButton color="inherit" aria-label="menu" sx={{ ml: 'auto' }}>
              <Menu />
            </IconButton>
          </Toolbar>
        </AppBar>
      </Box>
      <Container maxWidth="sm">{props.children}</Container>
    </React.Fragment>
  )
}

export default DefaultLayout
