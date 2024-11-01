import Filter from './components/filter.vue.js'
import History from './components/history.vue.js'
import Footer from './components/footer.vue.js'

const {ref} = Vue

class SearchState {
  to = ""
  subject = ""
  total = 0
  page = 0
  messages = []
  maxRows = 20
}

export default {
  data() {
    return {
      state: new SearchState(),
    }
  },
  methods: {
    deleteEvent(id) {
      console.log("deleting event id:", id)
    },
    filter() {
      console.log("Filter using ", this.state.to, " and ", this.state.subject)

      fetch(
        "/messages?" + new URLSearchParams(
          {
            "to": this.state.to,
            "subject": this.state.subject,
            "page": this.state.page,
            "rows": this.state.maxRows,
          }
        )
      )
        .then(response => response.json())
        .then(json => {
          this.state.messages = json.messages
          this.state.total = json.total
        })
        .catch(err => console.error(err))
    },
  },
  components: {Filter, History, Footer},
  template: `
    <section class="section">
      <Filter :state="state" :filter-func="filter"/>
      <History :state="state" :filter-func="filter" :delete-event-func="deleteEvent"/>
      <Footer :state="state"/>

    </section>
  `
}
