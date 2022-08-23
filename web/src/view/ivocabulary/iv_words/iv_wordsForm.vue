<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="单词UUID:">
          <el-input v-model.number="formData.uuid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="英文单词:">
          <el-input v-model="formData.word" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="音标（英音）:">
          <el-input v-model="formData.phoneticEn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="音标（美音）:">
          <el-input v-model="formData.phoneticUs" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="词义:">
          <el-input v-model="formData.meanings" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="在哪些考试中使用:">
          <el-input v-model="formData.examInfo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标记:">
          <el-switch v-model="formData.mark" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="被标记次数:">
          <el-input v-model.number="formData.markCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="笔记:">
          <el-input v-model="formData.notes" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="读音文件路径（英音）:">
          <el-input v-model="formData.audioFpAm" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="读音文件路径（美音）:">
          <el-input v-model="formData.audioFpEn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="读音文件路径（tts）:">
          <el-input v-model="formData.audioFpTts" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'WordInfo'
}
</script>

<script setup>
import {
  createWordInfo,
  updateWordInfo,
  findWordInfo
} from '@/api/iv_words'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        uuid: 0,
        word: '',
        phoneticEn: '',
        phoneticUs: '',
        meanings: '',
        examInfo: '',
        mark: false,
        markCount: 0,
        notes: '',
        audioFpAm: '',
        audioFpEn: '',
        audioFpTts: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWordInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewords
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createWordInfo(formData.value)
          break
        case 'update':
          res = await updateWordInfo(formData.value)
          break
        default:
          res = await createWordInfo(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
