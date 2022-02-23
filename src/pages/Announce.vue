<template>
    <v-container>
        <v-card height="600px" width="1000px">
                    <v-simple-table>
                        <template>
                            <thead>
                                <tr>
                                    <th>
                                        更新日付
                                    </th>
                                    <th>
                                        更新内容
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td>2022/2/13</td>
                                    <td>α版リリース</td>
                                </tr>
                            </tbody>
                        </template>
                    </v-simple-table>
                </v-card>
    </v-container>
</template>

<script>
export default ({
    name: 'announce'
})
</script>
