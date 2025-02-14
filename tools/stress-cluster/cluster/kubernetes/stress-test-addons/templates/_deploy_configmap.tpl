{{ define "stress-test-addons.deploy-configmap" }}
apiVersion: v1
kind: ConfigMap
metadata:
  name: "{{ .Release.Name }}-{{ .Release.Revision }}-test-resources"
  namespace: {{ .Release.Namespace }}
data:
  template: |
    {{- .Files.Get "stress-test-resources.json" | nindent 4 }}
{{ end }}
