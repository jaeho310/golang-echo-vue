import Router from 'vue-router'
import Home from '@/views/Home'
import Test from '@/views/Test'
import Error from '@/views/Error'
import Vue from 'vue'

Vue.use(Router)
export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: "Home",
            component: Home
        }
        ,{
            path: '/test',
            name: 'Test',
            component: Test
        },
        {
            path: '*',
            name: 'Error',
            component: Error
        }
    ]
})
