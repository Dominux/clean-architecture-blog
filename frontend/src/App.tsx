import React from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import {
  AppBar,
  Box,
  Toolbar,
  Typography,
  IconButton,
  Container,
} from '@mui/material'
import { Menu } from '@mui/icons-material'
import axios from 'axios'

import Home from './pages'

axios.defaults.baseURL = `${process.env.PROTOCOL || 'http'}://${
  process.env.HOST || 'localhost'
}:${process.env.PORT || 8080}`

export default function App() {
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

      <Container maxWidth="sm">
        {/* Webpage */}
        <Router>
          <Switch>
            <Route path="/">
              <Home />
            </Route>
          </Switch>
        </Router>
      </Container>
    </React.Fragment>
  )
}
