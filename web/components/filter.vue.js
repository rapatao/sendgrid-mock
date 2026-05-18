export default {
  props: {
    state: Object,
    filterFunc: Function,
  },
  data() {
    return {
      params: {
        to: "",
        subject: "",
      }
    }
  },
  methods: {
    clear() {
      this.params.to = ""
      this.params.subject = ""
      this.filter()
    },
    filter() {
      this.state.page = 0
      this.state.to = this.params.to
      this.state.subject = this.params.subject

      this.filterFunc()
    },
  },
  template: `
    <div class="box py-4">
      <div class="columns">
        <div class="column">
          <div class="field">
            <label for="to" class="label">To</label>
            <div class="control has-icons-left">
              <input id="to" class="input" type="email" placeholder="recipient@example.com"
                     v-model="params.to" v-on:keyup.enter="filter"/>
              <span class="icon is-small is-left">
                <i class="fas fa-envelope"></i>
              </span>
            </div>
          </div>
        </div>

        <div class="column">
          <div class="field">
            <label for="subject" class="label">Subject</label>
            <div class="control has-icons-left">
              <input id="subject" class="input" type="text" placeholder="Search subject..."
                     v-model="params.subject" v-on:keyup.enter="filter"/>
              <span class="icon is-small is-left">
                <i class="fas fa-heading"></i>
              </span>
            </div>
          </div>
        </div>

        <div class="column is-narrow">
          <div class="field">
            <label class="label">&nbsp;</label>
            <div class="control is-flex is-align-items-center" style="height: 2.5rem;">
              <div class="buttons">
                <button class="button is-link is-small" @click="filter">
                  <span class="icon is-small"><i class="fas fa-search"></i></span>
                  <span>Filter</span>
                </button>
                <button class="button is-light is-small" @click="clear">
                  <span class="icon is-small"><i class="fas fa-redo"></i></span>
                  <span>Clear</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  `
}
