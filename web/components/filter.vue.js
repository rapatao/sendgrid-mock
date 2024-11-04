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
    <div class="container">

      <div class="columns box">
        <div class="column is-two-fifths">
          <div class="field is-horizontal">
            <div class="field-label is-normal">
              <label for="to" class="label">To</label>
            </div>
            <div class="field-body">
              <div class="field">
                <p class="control">
                  <input id="to" class="input is-normal" type="email" placeholder="example@example.com"
                         v-model="params.to" v-on:keyup.enter="filter"/>
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="column is-two-fifths">
          <div class="field is-horizontal">
            <div class="field-label is-normal">
              <label for="subject" class="label">Subject</label>
            </div>
            <div class="field-body">
              <div class="field">
                <p class="control">
                  <input id="subject" class="input is-normal" type="text" placeholder="Subject"
                         v-model="params.subject" v-on:keyup.enter="filter"/>
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="column is-one-fifth">
          <div class="field is-grouped">
            <div class="control">
              <button class="button is-success" @click="filter">Filter</button>
            </div>
            <div class="control">
              <button class="button is-info" @click="clear">Clear</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  `
}
