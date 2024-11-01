const {ref} = Vue

export default {
  props: {
    state: Object,
    filterFunc: Function,
  },
  setup() {
    return {
      clear: ref(function () {
        this.state.to = null
        this.state.subject = null

        this.filterFunc()
      }),
    }
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
                         v-model="state.to"/>
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="column is-two-fifths">
          <div class="field is-horizontal">
            <div class="field-label is-normal">
              <label for="subject" class="label">From</label>
            </div>
            <div class="field-body">
              <div class="field">
                <p class="control">
                  <input id="subject" class="input is-normal" type="text" placeholder="Subject"
                         v-model="state.subject"/>
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="column is-one-fifth">
          <div class="field is-grouped">
            <div class="control">
              <button class="button is-link" @click="filterFunc()">Filter</button>
            </div>
            <div class="control">
              <button class="button is-warning" @click="clear()">Clear</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  `
}
