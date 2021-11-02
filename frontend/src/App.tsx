import React from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import axios from 'axios'

import DefaultLayout from './layouts/default'
import Home from './pages'
import Post from './pages/post'
import NotFound from './pages/404'

axios.defaults.baseURL = `${process.env.PROTOCOL || 'http'}://${
  process.env.HOST || 'localhost'
}:${process.env.PORT || 8080}`

export default function App() {
  return (
    <Router>
      <Switch>
        <Route component={DefaultLayout}>
          <Route exact path="/">
            <Home />
          </Route>
          <Route path="/posts/:id">
            <Post />
          </Route>
        </Route>
        <Route component={NotFound}>
          <Route path="*" />
        </Route>
      </Switch>
    </Router>
  )
}
