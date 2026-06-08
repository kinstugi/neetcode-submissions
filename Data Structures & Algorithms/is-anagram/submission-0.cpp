class Solution {
public:
    bool isAnagram(string s, string t) {
        int mp[26] = {0};
        for(char ch: s)
        {
            mp[ch - 'a']++;
        }

        for(char ch: t)
        {
            mp[ch - 'a']--;
        }

        for (int i = 0; i < 26; i++)
        {
            if (mp[i]) return false;
        }
        return true;
    }
};
