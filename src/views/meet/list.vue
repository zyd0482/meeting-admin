<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.title" placeholder="Title" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-button class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
        Search
      </el-button>
      <router-link :to="'/meet/edit/'+row.id">
        <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit">
          Add
        </el-button>
      </router-link>
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column label="ID" prop="id" sortable="custom" align="center" width="80px">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="Type" class-name="status-col" width="100">
        <template slot-scope="{row}">
          <el-tag :type="row.type | typeFilter">
            {{ row.type }}
          </el-tag>
        </template>
      </el-table-column>

       <el-table-column label="Banner" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.banner }}</span>
        </template>
      </el-table-column>

      <el-table-column label="Title" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.title }}</span>
        </template>
      </el-table-column>

      <el-table-column label="StartAt" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.start_at }}</span>
        </template>
      </el-table-column>

      <el-table-column label="Place" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.place }}</span>
        </template>
      </el-table-column>

      <el-table-column label="Person" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.person }}</span>
        </template>
      </el-table-column>

      <el-table-column label="Fee" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.fee }}</span>
        </template>
      </el-table-column>

      <el-table-column label="Status" class-name="status-col" width="100">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusFilter">
            {{ row.status }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column label="Detail" width="100">
        <template slot-scope="{row}">
          <el-button type="primary" size="mini" @click="handlePreview(row)">
            查看介绍
          </el-button>
        </template>
      </el-table-column>

      <el-table-column label="Actions" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row}">
          <router-link :to="'/meet/edit/'+row.id">
            <el-button type="primary" size="small" icon="el-icon-edit">
              Edit
            </el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import { fetchList } from '@/api/meet'
import Pagination from '@/components/Pagination'

export default {
  name: 'ComplexTable',
  components: { Pagination },
  filters: {
    stateFilter(state) {
      const stateMap = {
        1: '有效',
        2: '无效'
      }
      return stateMap[state]
    },
    typeFilter(type) {
      const typeMap = {
        1: '普通活动',
        2: '定期活动'
      }
      return typeMap[type]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20,
        title: undefined,
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.listLoading = false
        this.list = response.data.items
        this.total = response.data.total
      })
    },
    handlePreview(row) {
      console.log(row.detail)
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    }
  }
}
</script>
