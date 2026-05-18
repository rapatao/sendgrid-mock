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
  loading = false
}

export default {
  data() {
    return {
      state: new SearchState(),
      isFilterVisible: true,
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
      this.state.loading = true
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
        .finally(() => {
          this.state.loading = false
        })
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
  mounted() {
    const observer = new IntersectionObserver(
      ([entry]) => {
        this.isFilterVisible = entry.isIntersecting
      },
      { threshold: 0 }
    )
    if (this.$refs.filterContainer) {
      observer.observe(this.$refs.filterContainer)
    }
  },
  beforeMount() {
    this.filter()
  },
  components: {Filter, History, Footer},
  template: `
    <div>
      <div ref="filterContainer" class="filter-wrapper">
        <div class="container">
          <Filter :state="state" :filter-func="filter"/>
        </div>
      </div>
      
      <button v-show="!isFilterVisible" class="button is-link is-rounded fab-filters" @click="scrollTop" title="Go to filters">
        <span class="icon">
          <i class="fas fa-search"></i>
        </span>
      </button>

      <div class="container mt-0">
        <History :state="state" :filter-func="filter" :delete-func="deleteEvent"/>
        <Footer :state="state" :delete-all-func="deleteAll"/>
      </div>
    </div>
  `
}
