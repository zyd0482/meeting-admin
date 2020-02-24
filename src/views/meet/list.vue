<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="240px" label="类型">
        <template slot-scope="{row}">
          <el-tag :type="row.type | typeFilter">
            {{ row.type }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column min-width="240px" label="名称">
        <template slot-scope="{row}">
          <span>{{ row.title }}</span>
        </template>
      </el-table-column>

      <!-- <el-table-column width="100px" label="头图">
        <template slot-scope="{row}">
          <img class="banner" :src="row.banner">
        </template>
      </el-table-column> -->
      <el-table-column width="180px" align="center" label="时间">
        <template slot-scope="scope">
          <span>{{ scope.row.start_at | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="100px" label="地点">
        <template slot-scope="{row}">
          <span>{{ row.place }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="100px" label="人数">
        <template slot-scope="{row}">
          <span>{{ row.person }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="100px" label="人均消费">
        <template slot-scope="{row}">
          <span>{{ row.fee }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="100px" label="状态">
        <template slot-scope="{row}">
          <el-tag :type="row.state | stateFilter">
            {{ row.state }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column align="center" label="" width="200">
        <template slot-scope="scope">
          <router-link :to="'/meet/detail/'+scope.row.id">
            <el-button type="primary" size="small" icon="el-icon-edit">
              查看
            </el-button>
          </router-link>
          <router-link :to="'/meet/edit/'+scope.row.id">
            <el-button type="primary" size="small" icon="el-icon-edit">
              编辑
            </el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.page_limit" @pagination="getList" />
  </div>
</template>

<script>
import { fetchList } from '@/api/meet'
import Pagination from '@/components/Pagination'

export default {
  name: '活动列表',
  components: { Pagination },
  filters: {
    typeFilter(type) {
      const typeMap = {
        1: '常规活动',
        2: '定期活动'
      }
      return typeMap[type]
    },
    stateFilter(state) {
      const stateMap = {
        1: '正常',
        2: '暂停'
      }
      return stateMap[state]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        page_limit: 20
      },
      dialogVisible: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      this.listQuery.page_start_at = this.listQuery.page * this.listQuery.page_limit;
      fetchList(this.listQuery).then(response => {
        this.list = response.data.list
        this.total = response.data.total
        this.listLoading = false
      })
    }
  }
}
</script>
