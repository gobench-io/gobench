import firebase from 'firebase/app'
import { notification } from 'antd'
import 'firebase/auth'
import 'firebase/database'

const firebaseConfig = {
  apiKey: 'AIzaSyBJVhr2WZshEGR7egcxoygQIphKOkKVIYQ',
  authDomain: 'sellpixels-7d5d4.firebaseapp.com',
  databaseURL: 'https://sellpixels-7d5d4.firebaseio.com',
  projectId: 'sellpixels-7d5d4',
  storageBucket: 'cleanui-72a42.appspot.com',
  messagingSenderId: '338219933237'
}

firebase.initializeApp(firebaseConfig)
export const firebaseAuth = firebase.auth()
export const firebaseDatabase = firebase.database()

export async function login (email, password) {
  return firebaseAuth
    .signInWithEmailAndPassword(email, password)
    .then(() => true)
    .catch(error => {
      notification.warning({
        message: error.code,
        description: error.message
      })
    })
}

export async function register (email, password, name) {
  return firebaseAuth
    .createUserWithEmailAndPassword(email, password)
    .then(response => {
      if (response.user) {
        const { uid } = response.user
        firebaseDatabase
          .ref('users')
          .child(uid)
          .set({
            role: 'admin',
            name
          })
      }
      return true
    })
    .catch(error => {
      notification.warning({
        message: error.code,
        description: error.message
      })
    })
}

export async function currentAccount () {
  let userLoaded = false
  function getCurrentUser (auth) {
    return new Promise((resolve, reject) => {
      if (userLoaded) {
        resolve(firebaseAuth.currentUser)
      }
      const unsubscribe = auth.onAuthStateChanged(user => {
        userLoaded = true
        unsubscribe()
        const getUserData = async () => {
          if (user) {
            const userFields = await firebaseDatabase
              .ref('users')
              .child(user.uid)
              .once('value')
              .then(snapshot => {
                return snapshot.val()
              })
            const mergedUser = Object.assign(user, {
              id: user.uid,
              name: userFields.name,
              role: userFields.role,
              avatar: user.photoUrl
            })
            return mergedUser
          }
          return user
        }
        resolve(getUserData())
      }, reject)
    })
  }
  return getCurrentUser(firebaseAuth)
}

export async function logout () {
  return firebaseAuth.signOut().then(() => true)
}
