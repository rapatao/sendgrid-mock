import Filter from './components/filter.vue.js'
import History from './components/history.vue.js'
import Footer from './components/footer.vue.js'

const {ref} = Vue

class SearchState {
  to = null
  subject = null
  total = 0
  messages = []
}

let state = new SearchState();

function filter() {
  console.log("Filter using ", state.to, " and ", state.subject)
}

function deleteEvent(id) {
  console.log("deleting event id:", id)
}

export default {
  setup() {
    return {
      state: ref(state),
      filterFunc: ref(filter),
      deleteEventFunc: ref(deleteEvent)
    }
  },
  components: {Filter, History, Footer},
  template: `
    <section class="section">
      <Filter :state="state" :filter-func="filterFunc"/>
      <History :state="state" :delete-event-func="deleteEventFunc"/>
      <Footer :state="state"/>

    </section>
  `
}
