class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        map<std::string, vector<std::string>> mp;
        for(const string &s: strs)
        {
            std::string key(s);
            std::sort(key.begin(), key.end());
            mp[key].push_back(s);
        }

        vector<vector<std::string>> ans;
        ans.reserve(10);
        for (const auto &item: mp)
        {
            ans.push_back(item.second);
        }
        return ans;
    }
};
