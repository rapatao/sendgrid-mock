import Filter from './components/filter.vue.js'
import History from './components/history.vue.js'
import Footer from './components/footer.vue.js'

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
    scrollTop() {
      window.scroll({
        top: 0,
        left: 0,
        behavior: 'smooth'
      })
    },
    filter() {
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
          
          this.scrollTop()
        })
        .catch(err => console.error(err))
    },
    deleteEvent(id) {
      fetch("/messages/" + id, {
        method: "DELETE",
      }).then(_ => this.filter())
        .catch(err => console.error(err))
    },
    deleteAll() {
      fetch("/messages", {
        method: "DELETE",
      }).then(_ => this.filter())
        .catch(err => console.error(err))
    }
  },
  beforeMount() {
    this.filter()
  },
  components: {Filter, History, Footer},
  template: `
    <section class="section">
      <Filter :state="state" :filter-func="filter"/>
      <History :state="state" :filter-func="filter" :delete-func="deleteEvent"/>
      <Footer :state="state" :delete-all-func="deleteAll"/>
    </section>
  `
}
