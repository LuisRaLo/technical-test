import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider } from 'react-router-dom'
import router from './routes'
import './index.css'
import { Provider } from 'react-redux'
import { AppStore } from './states/store'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <Provider store={AppStore}>
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>
  </Provider>
)
