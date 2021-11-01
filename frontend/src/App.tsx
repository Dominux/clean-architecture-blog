import React from 'react'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from 'react-router-dom'
import axios from 'axios'

import Home from './pages'
import DefaultLayout from './layouts/default'
import Post from './pages/post'

axios.defaults.baseURL = `${process.env.PROTOCOL || 'http'}://${
  process.env.HOST || 'localhost'
}:${process.env.PORT || 8080}`

export default function App() {
  return (
    <Router>
      <Switch>
        <Route exact path="/">
          <DefaultLayout>
            <Home />
          </DefaultLayout>
        </Route>
        <Route path="/posts/:id"><Post/></Route>
        <Route>
          <Redirect to="/" />
        </Route>
      </Switch>
    </Router>
  )
}
